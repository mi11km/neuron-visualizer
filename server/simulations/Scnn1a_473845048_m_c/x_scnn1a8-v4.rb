#!/usr/bin/ruby

OUTPUTFILE = "Scnn1a_473845048_m_c_ion.txt"
INPUTFILE = "Scnn1a_473845048_m_c.swc"

#        [dummy, soma,             axon,                basal dendrite,        apical dendrite]
CM     = [-1,    1.0,              1.0,                 2.12,                  2.12]
G_LEAK = [-1,    5.71880766722e-06*1000, 0.00045738760076499994*1000, 3.2393273274400003e-06*1000, 9.5861855476200007e-05*1000]
RA     = [-1,    0.138280*1000,    0.138280*1000,       0.138280*1000,         0.138280*1000]
#G_NATS = [-1,    0.0, 0.0, 0.0, 0.0]
#G_KV3  = [-1,    0.0, 0.0, 0.0, 0.0]
#G_KP   = [-1,    0.0, 0.0, 0.0, 0.0]
#G_IM   = [-1,    0.0, 0.0, 0.0, 0.0]
#G_NAP  = [-1,    0.0, 0.0, 0.0, 0.0]
#G_KT   = [-1,    0.0, 0.0, 0.0, 0.0]
#G_SK   = [-1,    0.0, 0.0, 0.0, 0.0]
#G_IH   = [-1,    0.0, 0.0, 0.0, 0.0]
G_NATA = [-1,    0.0, 0.0, 0.0, 0.0]
#G_KD   = [-1,    0.000001 * 1000, 0.0, 0.0, 0.0]
#G_IMV2 = [-1,    0.000001 * 1000, 0.0, 0.0, 0.0]
#G_KV2  = [-1,    0.000001 * 1000, 0.0, 0.0, 0.0]
G_KD   = [-1,    0.0, 0.0, 0.0, 0.0]
G_IMV2 = [-1,    0.0, 0.0, 0.0, 0.0]
G_KV2  = [-1,    0.0, 0.0, 0.0, 0.0]
#G_CAHVA = [-1,  0.0, 0.0, 0.0, 0.0]
#G_CALVA = [-1,  0.0, 0.0, 0.0, 0.0]
G_CAHVA= [-1,    0.00053599731839199991*1000, 0.0, 0.0, 0.0]
G_CALVA = [-1,   0.0070061294358100008*1000, 0.0, 0.0, 0.0]

G_NATS = [-1,    0.98228995892999993*1000, 0.0, 0.0, 0.0]

G_NAP  = [-1,    0.000209348990528*1000    , 0.0, 0.0, 0.0]
G_KT   = [-1,    0.00073160714529799998*1000, 0.0, 0.0, 0.0]
G_KP   = [-1,    0.051758360920800002*1000, 0.0, 0.0, 0.0]
G_KV3  = [-1,    0.057264803402699994*1000, 0.0, 0.0, 0.0]

#G_SK   = [-1,    0.00019222004878899999*1000, 0.000001 * 1000, 0.000001 * 1000, 0.000001*1000] # 10^(-6) is correct for SK
G_SK   = [-1,    0.00019222004878899999*1000, 0.0, 0.0, 0.0] # 10^(-6) is correct for SK
G_IH   = [-1,    4.12225901169e-05*1000,     0.0,   0.0, 0.0]
G_IM   = [-1,    0.0012021154978800002*1000, 0.0, 0.0, 0.0]

=begin
G_NATS = [-1,    0.98228995892999993*1000, 0.0, 0.0, 0.0]
G_KV3  = [-1,    0.057264803402699994*1000, 0.0, 0.0, 0.0]
G_KP   = [-1,    0.051758360920800002*1000, 0.0, 0.0, 0.0]
G_IM   = [-1,    0.0012021154978800002*1000, 0.0, 0.0, 0.0]
G_NAP  = [-1,    0.000209348990528*1000    , 0.0, 0.0, 0.0]
G_KT   = [-1,    0.00073160714529799998*1000, 0.0, 0.0, 0.0]
G_SK   = [-1,    0.00019222004878899999*1000, 0.000001 * 1000, 0.000001 * 1000, 0.000001*1000] # 10^(-6) is correct for SK
G_IH   = [-1,    4.12225901169e-05*1000,     0.0,   0.0, 0.0]
G_NATA = [-1,    0.00001*1000,  0.0,  0.0,  0.0]
G_KD   = [-1,    0.0,  0.0,  0.0,  0.0]
G_IMV2 = [-1,    0.0,  0.0,  0.0,  0.0]
G_KV2  = [-1,    0.00001*1000,  0.00001*1000,  0.00001*1000,  0.00001*1000]
G_CAHVA= [-1,    0.00053599731839199991*1000, 0.0, 0.0, 0.0]
G_CALVA = [-1,   0.0070061294358100008*1000, 0.0, 0.0, 0.0]
=end
                                     
#GAMMA  = [-1,    0.0012510775510599999, 0.05, 0.05, 0.05 ]
#DECAY  = [-1,    717.91660042899991, 80.0, 80.0, 80.0 ]


Item = Struct.new(:type, :x, :y, :z, :r, :pid)

def main

  data = Array.new
  # swc format
  # id type x_pos y_pos z_pos(mum) radius(mum), parent_id
  IO.foreach(INPUTFILE){|l|
    id, type, x, y, z, r, pid = l.chomp.split
    data[id.to_i] = Item.new(type.to_i, x.to_f, y.to_f, z.to_f, r.to_f, pid.to_i)
  }

  open(OUTPUTFILE, "w"){|o|
    data.each_with_index{|item, id|
      type = item.type
      x, y, z, r, pid = item.x, item.y, item.z, item.r, item.pid
      x_par, y_par, z_par = data[pid].x, data[pid].y, data[pid].z
      # length = 2 * r
      length = if id == 0
                 2 * r # equal to diameter
               else
                 Math.sqrt( (x-x_par)**2 + (y-y_par)**2 + (z-z_par)**2 )
               end
      o.puts "#{G_LEAK[type]} #{G_NATS[type]} #{G_NATA[type]} #{G_NAP[type]} #{G_KV2[type]} #{G_KV3[type]} #{G_KP[type]} #{G_KT[type]} #{G_KD[type]} #{G_IM[type]} #{G_IMV2[type]} #{G_IH[type]} #{G_SK[type]} #{G_CAHVA[type]} #{G_CALVA[type]}"
    }
  }
end

main if __FILE__ == $0