var Unit = {
    edit: function (obj,id,title) {
        $('#myModal').modal()                      // 以默认值初始化
        $('#myModal').modal({ keyboard: false })   // initialized with no keyboard
        $('input[name="Title"]').val(title)
        var html = '<input type="hidden" name="Id" value="'+id+'"> '
        $('input[name="Title"]').parent().append(html)

        $('#myModal').modal('show')           // 初始化后立即调用 show 方法
    }
}