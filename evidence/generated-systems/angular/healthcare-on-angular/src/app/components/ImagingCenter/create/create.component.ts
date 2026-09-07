import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ImagingCenterService } from '../../../services/ImagingCenter.service';
import { ImagingCenter } from '../../../models/ImagingCenter';
import { SubBaseComponent } from '../../ImagingCenter/sub.base.component';

@Component({
    selector: 'app-create-imagingCenter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateImagingCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Add ImagingCenter';

    imagingCenterForm: FormGroup;
    imagingCenter: ImagingCenter;

    constructor( http: HttpClient,
        private imagingCenterService: ImagingCenterService,
        private fb: FormBuilder,
        private router: Router
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

    
    addImagingCenter(name, Facility, ImagingOrders, ImagingReports): void {
        this.imagingCenterService
        .addImagingCenter(name, Facility, ImagingOrders, ImagingReports)
            .subscribe(() => {
                this.router.navigate(['/indexImagingCenter']);
            });
    }

    ngOnInit(): void {
    }
}