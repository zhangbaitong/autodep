<?php
  $image_name='test13';
  $dockerfile_directory='/data/DockerManager';
  $a = exec('/usr/bin/docker  build -t '.$image_name.'  '.$dockerfile_directory,$out,$status);
  echo "---1----"; print_r($a);
  echo "---2----"; print_r($out);
  echo "---3----"; print_r($status);

?>
