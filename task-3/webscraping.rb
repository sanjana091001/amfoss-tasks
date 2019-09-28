require 'nokogiri'
require 'open-uri'
document = Nokogiri::HTML(open('https://www.google.com/search?hl=en&source=hp&ei=p6uMXcqsBeLez7sPpaiA0AU&q=linux&oq=linux&gs_l=psy-ab.3..0l10.911.2075..2511...0.0..0.421.1054.0j2j0j1j1......0....1..gws-wiz.......0i131.kfth2lBmdfo&ved=0ahUKEwiKnNmxu-7kAhVi73MBHSUUAFoQ4dUDCAU&uact=5').read)
data = document.xpath('//div//a/div')
data.each do |data|
puts(data.text)
end
