/**
 * Created by rick on 2017/6/19.
 */
/*var Vue = require('vue');
var VueResource = require('vue-resource.min.js');
Vue.use(VueResource);*/

var vm = new Vue({
    el: '#urlportal',
    data: {
        portaldata:[]
    },
    created:function () {
        this.getportaldata()
    },
    methods:{
        getportaldata:function () {
        this.$http.get('/api/v1/infos.do').then(res => {
            this.portaldata = res.body.data;
            },res =>{

        });
    }
    }

});

