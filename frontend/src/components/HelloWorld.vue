<script lang="ts" setup>
import {reactive, ref, watch} from 'vue'
import {Greet, ScreenRecord} from '../../wailsjs/go/main/App'
import {
  ArrowLeft,
  ArrowRight,
  ArrowDown,
  Setting,
  Document,
  Delete,
  Edit,
  Share,
} from '@element-plus/icons-vue'
import CanvasView from "./video/CanvasView.vue";

const isCollapse = ref(true)
const collapseButtonClass = ref("aside-button-coll")
const elIconRight = ref("el-icon--right")
const activeIndex = ref('1')
const input = ref('')

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})





function changeCollapse() {
  isCollapse.value = !isCollapse.value
}

function greet() {

  Greet(data.name).then(result => {
    data.resultText = result
  })
}


function screenRecord() {
  ScreenRecord()
}

watch(
    isCollapse,
    (oval, val) => {
      if (!val) {
        collapseButtonClass.value = "aside-button-coll"
        elIconRight.value = "el-icon--right"
      } else {
        collapseButtonClass.value = "aside-button"
        elIconRight.value = "el-icon--right-coll"
      }
    })

</script>

<template>
  <main>
    <div class="common-layout">
      <el-container>
        <el-header class="head">

            <div style="position:absolute; top:40%;">
              LOGO
            </div>

          <el-input v-model="input" placeholder="Please input" class="head-input" />
            <el-dropdown style="position: absolute; top: 20%; left: 90%; margin-right: 5%">
              <el-button type="text" style="color: white;">
                Dropdown List
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>Action 1</el-dropdown-item>
                  <el-dropdown-item>Action 2</el-dropdown-item>
                  <el-dropdown-item>Action 3</el-dropdown-item>
                  <el-dropdown-item disabled>Action 4</el-dropdown-item>
                  <el-dropdown-item divided>Action 5</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>

        </el-header>
        <el-container>
          <el-aside class="aside">
            <!--          <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">-->
            <!--            <el-radio-button :label="false">expand</el-radio-button>-->
            <!--            <el-radio-button :label="true">collapse</el-radio-button>-->
            <!--          </el-radio-group>-->
            <el-menu
                default-active="2"
                class="el-menu-vertical-demo"
                :collapse="isCollapse"
                @open="handleOpen"
                @close="handleClose"
                active-text-color="#ffd04b"
                background-color="#73767a"
                text-color="#fff"
            >

              <el-menu-item index="2">
                <el-icon>
                  <Edit/>
                </el-icon>
                <template #title>屏幕录制</template>
              </el-menu-item>
              <el-menu-item index="3">
                <el-icon>
                  <document/>
                </el-icon>
                <template #title>直播推流</template>
              </el-menu-item>

              <el-menu-item index="4">
                <el-icon>
                  <setting/>
                </el-icon>
                <template #title>视频转码</template>
              </el-menu-item>
              <el-menu-item index="5">
                <el-icon>
                  <setting/>
                </el-icon>
                <template #title>设置</template>
              </el-menu-item>
            </el-menu>
            <button :class=collapseButtonClass @click="changeCollapse">
              <el-icon :class=elIconRight>
                <ArrowRight/>
              </el-icon>
            </button>
          </el-aside>
          <el-main>
<!--            <div id="result" class="result">{{ data.resultText }}</div>-->
<!--            <div id="input" class="input-box">-->
<!--              <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>-->
<!--              <button class="btn" @click="screenRecord">Greet</button>-->
<!--            </div>-->
            <canvas-view></canvas-view>
<!--            <video src="/Users/private/wails_test/myproject/230629104241.mp4" width="720" height="480" controls preload="auto"></video>-->
          </el-main>
        </el-container>
      </el-container>
    </div>
  </main>
</template>

<style scoped>

.aside {
  width: fit-content;
  overflow: hidden
}

.head {
  background-color: grey;
  position:relative;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.el-menu-vertical-demo {
  /*height: 92.15vh;*/
  height: 87vh;
  background: rgba(255, 255, 255, 0.1);
  border: #1b2636;
  /*width: 5vw;*/
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  /*width: 200px;*/
  /*min-height: 400px;*/
  /*height: 92.15vh;*/
  width: 200px;
  background: rgba(255, 255, 255, 0.1);
}

.aside-button {
  height: 5vh;
  width: 200px;
  margin-left: 0px;
  display: flex;
  transition-property: width;
  transition-duration: 0.435s;
}

.aside-button-coll {
  display: flex;
  height: 5vh;
  width: 64px;
  transition-property: width;
  transition-duration: 0.435s;
}

.el-icon--right {
  margin: 2vh 20px;
}

.el-icon--right-coll {
  margin: 2vh 90px;
}

.el-menu-demo {
  width: 99vw;
  margin-left: 0px;
  background: rgba(255, 255, 255, 0);
  border: hidden;
  position:relative;
}

.flex-grow {
  flex-grow: 1;
}


.el-dropdown-link {
  cursor: pointer;
  color: #b3e19d;
  display: flex;
  align-items: center;
  border: hidden;
}

.head-input {
  width: 20%;
  position: absolute;
  top: 20%;
  left: 40%;
}
</style>
