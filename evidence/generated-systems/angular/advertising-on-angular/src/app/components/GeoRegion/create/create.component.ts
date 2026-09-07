import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { GeoRegionService } from '../../../services/GeoRegion.service';
import { GeoRegion } from '../../../models/GeoRegion';
import { SubBaseComponent } from '../../GeoRegion/sub.base.component';

@Component({
    selector: 'app-create-geoRegion',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateGeoRegionComponent extends SubBaseComponent implements OnInit {

    title = 'Add GeoRegion';

    geoRegionForm: FormGroup;
    geoRegion: GeoRegion;

    constructor( http: HttpClient,
        private geoRegionService: GeoRegionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.geoRegionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      name: ['', Validators.required],
      Parent: ['', ],
      Children: ['', ],
      RegionType: ['', ]
        });
    }

    
    addGeoRegion(code, name, Parent, Children, RegionType): void {
        this.geoRegionService
        .addGeoRegion(code, name, Parent, Children, RegionType)
            .subscribe(() => {
                this.router.navigate(['/indexGeoRegion']);
            });
    }

    ngOnInit(): void {
    }
}