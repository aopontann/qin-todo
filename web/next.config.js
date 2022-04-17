/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: ['assets.example.com', 'placehold.jp', 'lh3.googleusercontent.com'],
  },
};

module.exports = nextConfig;
