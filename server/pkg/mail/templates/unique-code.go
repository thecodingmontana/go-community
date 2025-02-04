package templates

func UniqueCodeTemplate() string {
	return `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Template</title>
    <style>
        /* Inlining Tailwind CSS styles here (you would typically use a tool to do this) */
        @import url('https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css');
    </style>
</head>
<body class="bg-gray-50">
    <div class="max-w-2xl mx-auto bg-white p-6 rounded-lg shadow-lg mt-12">
        <header class="text-center mb-6">
            <h1 class="text-3xl font-bold text-indigo-600">Hello, User!</h1>
        </header>
        <main>
            <p class="text-lg text-gray-700 mb-4">
                We’re excited to have you on board! Here's a quick guide on getting started.
            </p>
            <div class="bg-gray-100 p-4 rounded-lg mb-6">
                <h2 class="text-2xl font-semibold text-indigo-600">Step 1: Get Started</h2>
                <p class="text-gray-600 mt-2">
                    Visit our dashboard to explore features and configure your account.
                </p>
            </div>
            <div class="bg-gray-100 p-4 rounded-lg">
                <h2 class="text-2xl font-semibold text-indigo-600">Step 2: Explore</h2>
                <p class="text-gray-600 mt-2">
                    Discover powerful tools to enhance your workflow. Start with our tutorial.
                </p>
            </div>
        </main>
        <footer class="text-center mt-6">
            <p class="text-sm text-gray-600">
                Best regards,<br />
                The Example Team
            </p>
        </footer>
    </div>
</body>
</html>
	`
}
