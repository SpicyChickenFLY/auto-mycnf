$(document).ready(function () {
    function genText() {
        var result = "";
        // alert("OK");
        $(".param").each(
            function (key, value) {
                // $("#output").append($(this).attr("data-type")+ "<br/>");
                if ($(this).attr("data-type") == "chapter") { // 分组
                    result += $(this).text() + "\n";
                } else  { // 参数
                    // 判断是否被激活
                    if (!$(this).children("input[type='checkbox']").is(":checked")) {
                        result += "# ";
                    }
                    // 判断参数类型
                    if ($(this).attr("data-type") == "single") {
                        result += $(this).children("label").text() + "\n";
                    } else {
                        result += 
                            $(this).children("label").text() + "=" + 
                            $(this).children("input[type='text']").attr("value") +
                            "\n";
                    }
                }
            }
        )
        return result;
    }

    $("button#generate").click(function () {
        $.ajax({
            url:"http://localhost:8080/mycnf/file/gen",
            type: "post",
            data: {
                text: genText()
            },
            success:function(result) {
		    	alert(result);
            },
            error: function(result) {
                alert(result);
            }
        });
    });


});

