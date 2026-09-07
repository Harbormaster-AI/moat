import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { GeoRegionService } from '../../../services/GeoRegion.service';
import { SubBaseComponent } from '../../GeoRegion/sub.base.component';


@Component({
    selector: 'app-edit-geoRegion',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditGeoRegionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit GeoRegion';

    geoRegionForm: FormGroup;
    geoRegion: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: GeoRegionService,
        private fb: FormBuilder
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

    
    updateGeoRegion(code, name, Parent, Children, RegionType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateGeoRegion(code, name, Parent, Children, RegionType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexGeoRegion']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getGeoRegion(params['id']).subscribe(res => {
                this.geoRegion = res;
            });
        });
    }
}