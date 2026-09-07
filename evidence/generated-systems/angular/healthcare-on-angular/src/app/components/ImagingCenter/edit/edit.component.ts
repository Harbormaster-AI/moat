import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ImagingCenterService } from '../../../services/ImagingCenter.service';
import { SubBaseComponent } from '../../ImagingCenter/sub.base.component';


@Component({
    selector: 'app-edit-imagingCenter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditImagingCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ImagingCenter';

    imagingCenterForm: FormGroup;
    imagingCenter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ImagingCenterService,
        private fb: FormBuilder
) {
        super(http);
        this.imagingCenterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Facility: ['', ],
      ImagingOrders: ['', ],
      ImagingReports: ['', ]
        });
    }

    
    updateImagingCenter(name, Facility, ImagingOrders, ImagingReports): void {
        this.route.params.subscribe((params) => {

                        this.service.updateImagingCenter(name, Facility, ImagingOrders, ImagingReports, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexImagingCenter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getImagingCenter(params['id']).subscribe(res => {
                this.imagingCenter = res;
            });
        });
    }
}