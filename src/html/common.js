/**
 * Created by zhoulf on 2018/3/22.
 */

function HashMap() {
    var obj = new Object();
    this.put = function(key,value) {
        obj[key] = value;
    } ;
    this.get = function (key) {
        if (obj.hasOwnProperty(key)) {
            return obj[key];
        } else {
            return null;
        }
    };
    this.contains = function(key){
        return obj.hasOwnProperty(key)
    };
    this.remove = function(key) {
        if (obj.hasOwnProperty(key)) {
            var value = obj[key];
            delete obj[key];
            return value;
        } else {
            return null;
        }
    } ;
    this.size = function(){
        return Object.getOwnPropertyNames(obj).length;
    } ;
    this.keys = function() {
        return Object.getOwnPropertyNames(obj);
    };
    this.removeAll = function() {
        var properties = this.keys();
        for (var i = 0 ;i < properties.length;i++) {
            this.remove(properties[i]);
        }
    };


}

function HashSet(){
    var present = new Object();
    var hashMap = new HashMap();
    this.put = function(key) {
        if (hashMap.contains(key)) {
            return false;
        } else {
            hashMap.put(key, present);
            return true;
        }
    };
    this.contains = function(key) {
        return hashMap.contains(key);
    };
    this.remove = function(key) {
        hashMap.remove(key);
    };
    this.size = function() {
        return hashMap.size();
    };
    this.getAll = function(){
        return hashMap.keys();
    };
    this.removeAll = function() {
        hashMap.removeAll();
    };

}
