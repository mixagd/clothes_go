import {Component, Input, OnInit} from '@angular/core';
import {Cloth} from '../model';

@Component({
  selector: 'app-cloth-item',
  templateUrl: './cloth-item.component.html',
  styleUrls: ['./cloth-item.component.css']
})
export class ClothItemComponent {

  @Input()
  cloth: Cloth;

}
