$(function() {
  $("form").on("submit", function(e) {
    e.preventDefault();
    var md = $("textarea").val();
    var x = $.post("/markdown", md);
    x.done(function(html) {
      var el = $("#preview");
      el.html(html);
      el.find("pre code").each(function(i, block) {
        hljs.highlightBlock(block);
      });
    });

    x.fail(function(res) {
      alert(res.responseText);
    });
  });
});
