package main

import (
    "html/template"
    "io"
    "log"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// Template renderer struct
type TemplateRenderer struct {
    templates *template.Template
}

// Render implements echo.Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = renderer

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())
    e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
        XSSProtection:         "1; mode=block",
        ContentTypeNosniff:    "nosniff",
        XFrameOptions:         "SAMEORIGIN",
        HSTSMaxAge:           3600,
        ContentSecurityPolicy: "default-src 'self' unpkg.com",
    }))

    e.Static("/static", "static")

    // Routes
    setupRoutes(e)

    // Start server
    log.Fatal(e.Start(":3000"))
}

func setupRoutes(e *echo.Echo) {

    e.GET("/", handleHome)
    e.GET("/main-content", handleMainContent)
    e.GET("/shop", handleShop)
    e.GET("/about", handleAbout)
    e.GET("/cart", handleCart)
    
    // Product routes
    e.GET("/product/:id", handleProductDetail)
    
    // Cart routes
    e.POST("/cart/add", handleAddToCart)
    e.POST("/cart/remove", handleRemoveFromCart)
    e.POST("/cart/update", handleUpdateCart)
}

// Route handlers
func handleHome(c echo.Context) error {
    return c.Render(200, "index.html", nil)
}

func handleMainContent(c echo.Context) error {
    return c.HTML(200, `
        <div class="container mx-auto px-4 py-8">
            <nav class="flex justify-between items-center mb-8">
                <img src="/static/images/logo.svg" alt="PurePower Athletics" class="h-12">
                <div class="space-x-4">
                    <a href="/shop" class="text-gray-800 hover:text-gray-600">Shop</a>
                    <a href="/about" class="text-gray-800 hover:text-gray-600">About</a>
                    <a href="/cart" class="text-gray-800 hover:text-gray-600">Cart</a>
                </div>
            </nav>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                <!-- Featured products will go here -->
            </div>
        </div>
    `)
}

func handleShop(c echo.Context) error {
    // TODO: Implement shop page with products
    return c.Render(200, "shop.html", map[string]interface{}{
        "title": "Shop - PurePower Athletics",
    })
}

func handleAbout(c echo.Context) error {
    return c.Render(200, "about.html", map[string]interface{}{
        "title": "About - PurePower Athletics",
    })
}

func handleCart(c echo.Context) error {
    return c.Render(200, "cart.html", map[string]interface{}{
        "title": "Cart - PurePower Athletics",
    })
}

func handleProductDetail(c echo.Context) error {
    productID := c.Param("id")
    // TODO: Fetch product details from database
    return c.Render(200, "product.html", map[string]interface{}{
        "title":     "Product Detail - PurePower Athletics",
        "productID": productID,
    })
}

// Cart handlers
func handleAddToCart(c echo.Context) error {
    // TODO: Implement add to cart functionality
    return c.JSON(200, map[string]string{
        "message": "Product added to cart",
    })
}

func handleRemoveFromCart(c echo.Context) error {
    // TODO: Implement remove from cart functionality
    return c.JSON(200, map[string]string{
        "message": "Product removed from cart",
    })
}

func handleUpdateCart(c echo.Context) error {
    // TODO: Implement update cart functionality
    return c.JSON(200, map[string]string{
        "message": "Cart updated",
    })
}
