<template>
    <!-- <Header /> -->
    <div class="container">
        <div class="search">
            <h1>网签备案</h1>

        </div>
        <div class="search">
            <!-- <p style="margin-left: 10px">区域:</p>
            <el-select v-model="state.area" class="m-2" placeholder="Select" size="large" @change="loadData()">
                <el-option v-for="item in state.areaOptions" :key="item.value" :label="item.label"
                    :value="item.value" />
            </el-select> -->
            <p style="margin-left: 10px">最近时间:</p>
            <el-select v-model="state.limit" class="m-2" placeholder="Select" size="large" @change="loadData()">
                <el-option v-for="item in state.limitOptions" :key="item.value" :label="item.label"
                    :value="item.value" />
            </el-select>
            <!-- <span style="margin-left: 10px">按:</span> -->
            <el-select v-model="state.group" class="m-2" placeholder="Select" size="large" @change="loadData()">
                <el-option v-for="item in state.groupOptions" :key="item.value" :label="item.label"
                    :value="item.value" />
            </el-select>
        </div>
        <div>
            <!-- <h2>{{ state.area }}</h2> -->
            <el-card class="user-views-card" shadow="hover">
                <div class="account-growth" ref="accountGrowthChartRef"></div>
            </el-card>
        </div>
    </div>
</template>

<script setup lang="ts">
// import indexCover from '/@/assets/index/index-cover.svg'
// import { useSiteConfig } from '/@/stores/siteConfig'
// import { useMemberCenter } from '/@/stores/memberCenter'
import Header from '/@/layouts/frontend/components/header.vue'
import Footer from '/@/layouts/frontend/components/footer.vue'
import { houseGet } from '/@/api/frontend/index'
import { ref, reactive, onMounted, onBeforeMount, nextTick, onActivated } from 'vue'
import * as echarts from 'echarts'

const value = ref('')

const accountGrowthChartRef = ref<HTMLElement>()
const state: {
    area: string,
    limit: any,
    group: any,
    data: any,
    xData: any[],
    yData: any[],
    sData: any[],
    charts: any[],
    areaOptions: any[],
    limitOptions: any[],
    groupOptions: any[],
} = reactive({
    area: '闽侯',
    limit: 30,
    group: 'day',
    data: {},
    charts: [],
    xData: [],
    yData: [],
    sData: [],
    areaOptions: [],
    limitOptions: [],
    groupOptions: [],
})

const echartsResize = () => {
    nextTick(() => {
        for (const key in state.charts) {
            state.charts[key].resize()
        }
    })
}
const initUserGrowthChart = () => {
    const userGrowthChart = echarts.init(accountGrowthChartRef.value!)
    const option = {
        tooltip: {
            trigger: 'axis'
        },
        legend: {
            data: state.areaOptions
        },
        xAxis: {
            type: 'category',
            data: state.xData,
        },
        yAxis: {
            type: 'value'
        },

        series: state.sData
    }
    userGrowthChart.setOption(option)
    state.charts.push(userGrowthChart)
}
// const siteConfig = useSiteConfig()
// const memberCenter = useMemberCenter()

const loadData = () => {
    houseGet(state.area, state.limit,state.group).then((res) => {
        if (res.data) {
            console.log(res.data.xData)
            // console.log(res.data.yData)
            // console.log(res.data.areaRangeData)
            console.log(res.data.seriesData)
            state.xData = res.data.xData
            // state.yData = res.data.yData
            state.areaOptions = res.data.areaRangeData
            state.limitOptions = res.data.limitRange
            state.groupOptions = res.data.groupData
            state.sData = res.data.seriesData
            initUserGrowthChart()
        }
    }).catch((err: any) => {
        console.log(err)
    })
}


onMounted(() => {
    loadData()
    window.addEventListener('resize', echartsResize)
})

onBeforeMount(() => {
    for (const key in state.charts) {
        state.charts[key].dispose()
    }
    window.removeEventListener('resize', echartsResize)
})

onActivated(() => {
    echartsResize()
})
</script>

<style scoped lang="scss">
.container {
    width: 100vw;
    height: 100vh;
    background: url(/@/assets/bg.jpg) repeat;
    color: var(--el-color-white);

    .main {
        height: calc(100vh - 120px);
        padding: 0;

        .main-container {
            display: flex;
            height: 100%;
            width: 66%;
            margin: 0 auto;
            align-items: center;
            justify-content: space-between;

            .main-left {
                padding-right: 50px;

                .main-title {
                    font-size: 45px;
                }

                .main-content {
                    padding-top: 20px;
                    padding-bottom: 40px;
                    font-size: var(--el-font-size-large);
                }
            }

            .main-right {
                img {
                    width: 380px;
                }
            }
        }
    }
}

.header {
    background-color: transparent !important;
    box-shadow: none !important;
    position: fixed;
    width: 100%;

    :deep(.header-logo) {
        span {
            padding-left: 4px;
            color: var(--el-color-white);
        }
    }

    :deep(.frontend-header-menu) {
        background: transparent;

        .el-menu-item,
        .el-sub-menu .el-sub-menu__title {
            color: var(--el-color-white);

            &.is-active {
                color: var(--el-color-white) !important;
            }

            &:hover {
                background-color: transparent;
                color: var(--el-menu-hover-text-color);
            }
        }
    }
}

.footer {
    color: var(--el-text-color-secondary);
    background-color: transparent !important;
    position: fixed;
    bottom: 0;
}

@media screen and (max-width: 1024px) {
    .main-container {
        width: 90% !important;
        flex-wrap: wrap;
        align-content: center;
        justify-content: center !important;

        .main-right {
            padding-top: 50px;
        }
    }
}

@media screen and (max-width: 375px) {
    .main-right img {
        width: 300px !important;
    }
}

.account-growth {
    width: 100%;
    height: 600px;
}

.search {
    width: 100%;
    height: 100px;
    display: flex;
    justify-content: center;
    align-items: center;

    .el-input {
        width: 300px;
    }
}
</style>
