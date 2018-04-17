import {Component, Input, OnInit} from '@angular/core';
import {Cloth} from '../model';
import {ClothesService} from '../clothes.service';
import {ActivatedRoute, Router} from '@angular/router';

export interface ClothFormData {
  type?: string;
  colour?: string;
  fit?: string;
  owner?: string;
}

@Component({
  selector: 'app-cloth-form',
  templateUrl: './cloth-form.component.html',
  styleUrls: ['./cloth-form.component.css']
})
export class ClothFormComponent implements OnInit {

  constructor(private clothesService: ClothesService, private router: Router, private route: ActivatedRoute) {}

  formData: ClothFormData;

  @Input()
  cloth: Cloth;

  ngOnInit() {
    if (this.isEdit()) {
      this.clothesService.getCloth(this.route.snapshot.paramMap.get('id')).subscribe(c => this.formData = c);
    }
    this.formData = {
      type: '',
      colour: '',
      fit: '',
      owner: ''
    };
  }

  createOrUpdateCloth(valid: boolean) {
    if (valid && !this.isEdit()) {
      this.clothesService.createNewCloth(this.formData).subscribe(() => this.router.navigate(['/clothes']));
    } else {
      this.clothesService.updateCloth(this.route.snapshot.paramMap.get('id'), this.formData)
        .subscribe(() => this.router.navigate(['/clothes']));
    }
  }

  cancel() {
    this.router.navigate(['/clothes']);
  }

  isEdit(): boolean {
    return !!this.route.snapshot.paramMap.get('id');
  }

}
