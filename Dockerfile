FROM alpine
ADD github.com/nirolee/product.git /github.com/nirolee/product.git
ENTRYPOINT [ "/github.com/nirolee/product.git" ]
