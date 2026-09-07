import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ExposureService } from '../../../services/Exposure.service';
import { SubBaseComponent } from '../../Exposure/sub.base.component';


@Component({
    selector: 'app-edit-exposure',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditExposureComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Exposure';

    exposureForm: FormGroup;
    exposure: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ExposureService,
        private fb: FormBuilder
) {
        super(http);
        this.exposureForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  Claim: ['', ],
      PolicyCoverage: ['', ],
      InsuredObject: ['', ],
      Reserves: ['', ],
      Payments: ['', ],
      ExposureType: ['', ],
      Status: ['', ]
        });
    }

    
    updateExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexExposure']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getExposure(params['id']).subscribe(res => {
                this.exposure = res;
            });
        });
    }
}