import createAxios from '/@/utils/axios'
import { useSiteConfig } from '/@/stores/siteConfig'
import { useMemberCenter } from '/@/stores/memberCenter'
import { setTitle } from '/@/utils/common'

const controllerUrl = '/api/index/'

const houseUrl = '/house/'

export function index() {
    const siteConfig = useSiteConfig()
    const memberCenter = useMemberCenter()

    if (siteConfig.site_name) {
        return
    }

    return createAxios({
        url: controllerUrl + 'index',
        method: 'get',
    }).then((res) => {
        setTitle(res.data.site.site_name)
        siteConfig.dataFill(res.data.site)
        memberCenter.setStatus(res.data.open_member_center)
        if (!res.data.open_member_center) memberCenter.setLayoutMode('Disable')
    })
}


export function houseGet(area: string, limit: string,group: string): ApiPromise {
    return createAxios({
        url: houseUrl + 'get',
        method: 'GET',
        params: {
            area: area,
            limit: limit,
            group: group,
        },
    }) as ApiPromise
}