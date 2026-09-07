import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsurancePayerService } from '../../../services/InsurancePayer.service';
import { SubBaseComponent } from '../../InsurancePayer/sub.base.component';


@Component({
    selector: 'app-edit-insurancePayer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsurancePayerComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InsurancePayer';

    insurancePayerForm: FormGroup;
    insurancePayer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsurancePayerService,
        private fb: FormBuilder
) {
        super(http);
        this.insurancePayerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      Plans: ['', ],
      Claims: ['', ],
      PayerType: ['', ]
        });
    }

    
    updateInsurancePayer(name, website, Plans, Claims, PayerType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsurancePayer(name, website, Plans, Claims, PayerType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsurancePayer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsurancePayer(params['id']).subscribe(res => {
                this.insurancePayer = res;
            });
        });
    }
}