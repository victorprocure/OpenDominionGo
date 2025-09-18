SELECT id, ip_address, isp, organization, country, region, city, vpn, score, data
FROM user_origin_lookups
WHERE ip_address = $1;
