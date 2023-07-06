<template>
  <div id="back">
    <video id="background_video" autoplay height="480" v-show="false"></video>
    <canvas id="background_canvas" height="480"></canvas>
  </div>
</template>

<script setup lang="ts">

import {onMounted,ref} from "vue";
import {JsLog} from "../../../wailsjs/go/main/App";


const back_canvas = ref<HTMLCanvasElement>()
const back_ctx = ref<CanvasRenderingContext2D>()
const isDragging = ref(false)
const offsetX = ref(0)
const offsetY = ref(0)
const lastMouseX = ref(0)
const lastMouseY = ref(0)

onMounted(()=> {

  back_canvas.value = initBackGroundCanvas()
  createCanvas1(back_canvas.value, 1080)
  // createCanvas2(canvas.value)

})

function initBackGroundCanvas():HTMLCanvasElement {

  const canvas = document.getElementById("background_canvas") as HTMLCanvasElement;
  const video = document.getElementById("background_video") as HTMLVideoElement;
  canvas.style.backgroundColor = "black"
  canvas.style.width = "80%"
  const ctx = canvas.getContext("2d") as CanvasRenderingContext2D;
  back_ctx.value = ctx
  back_ctx.value.drawImage(video, 0, 0)
  return canvas
}

function createCanvas1(canvas:HTMLCanvasElement, height:number) {
  let videoElement = getViewVideo()

  console.log("videoElement", videoElement)


  if (videoElement.parentNode !== null) {
    videoElement.parentNode.appendChild(canvas)
  }
  const h = 1080;
  const w = h * 16 / 9
  canvas.width = w;
  canvas.height = h;
  if (back_ctx.value != null) {
    let timer = setInterval(function(){
      // console.log("test")
      if (!isDragging.value ) {
        if(back_ctx.value != null) {
          back_ctx.value.drawImage(videoElement, offsetX.value,offsetY.value, w, h)
          back_ctx.value.drawImage(videoElement, 0, 0, 1280, 720)
        }

      }

    })
    canvas.addEventListener("mousedown", function (e:MouseEvent) {
      isDragging.value = true
    })
    canvas.addEventListener("mousemove", function (e:MouseEvent) {
      if (isDragging.value && back_ctx.value != null) {

        back_ctx.value.clearRect(0, 0, canvas.width, canvas.height)
        // let mouseX = e.offsetX
        // let mouseY = e.offsetY

        // let deltaX = (mouseX - lastMouseX.value)
        // let deltaY = (mouseY - lastMouseY.value)

        lastMouseX.value = offsetX.value
        lastMouseY.value = offsetY.value

        offsetX.value = e.offsetX
        offsetY.value =  e.offsetY

        back_ctx.value.drawImage(videoElement, offsetX.value, offsetY.value , w, h)
        back_ctx.value.drawImage(videoElement, 0, 0, 1280, 720)
      }
    })
    canvas.addEventListener("mouseup", function() {
      isDragging.value = false
    })
  }

  function createCanvas2(canvas:HTMLCanvasElement) {
    let videoElement = getViewVideo()
    const ctx = canvas.getContext("2d") as CanvasRenderingContext2D;
    if (videoElement.parentNode !== null) {
      videoElement.parentNode.appendChild(canvas)
    }
    const w = 1280;
    const h = 720;
    canvas.width = w;
    canvas.height = h;
    let timer = setInterval(function(){
      // console.log("test")
      ctx.drawImage(videoElement, 0,0, w, h)
    })
  }

}




function getViewVideo():HTMLVideoElement{
  const divElement = document.getElementById("videoElement")
  const videoElement = document.createElement( "video") as HTMLVideoElement;
  videoElement.autoplay = true
  videoElement.height= 480
  if (divElement != null) {
    divElement.appendChild(videoElement)
  }

  try {
    navigator.mediaDevices.getDisplayMedia({
      video: {  frameRate: {ideal: 30}}
    })
        .then(stream => {
          // 将视频流分配给视频元素

          if (videoElement != null) {
            videoElement.srcObject = stream;
            videoElement.style.display = "none"
            return videoElement
          }
        })
        .catch(error => {
          JsLog(error.toString())
        });
  } catch (e) {
    if (typeof e === "string") {
      JsLog(e.toUpperCase()) // works, `e` narrowed to string
    } else if (e instanceof Error) {
      JsLog(e.message) // works, `e` narrowed to Error
    }

  }
  return videoElement
}

function initCanvasAndVideo() {

  // video.value = new HTMLVideoElement(document.getElementById('video'));
  // c1 = document.getElementById('c1');
  // ctx1 = c1.getContext('2d');
  // c2 = document.getElementById('c2');
  // ctx2 = c2.getContext('2d');
  // video.addEventListener('play', function() {
  //   self.width = self.video.videoWidth / 2;
  //   self.height = self.video.videoHeight / 2;
  //   self.timerCallback();
  // }, false);
}
</script>

<style scoped>

</style>