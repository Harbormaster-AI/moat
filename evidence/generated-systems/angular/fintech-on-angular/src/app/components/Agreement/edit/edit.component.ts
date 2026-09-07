import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AgreementService } from '../../../services/Agreement.service';
import { SubBaseComponent } from '../../Agreement/sub.base.component';


@Component({
    selector: 'app-edit-agreement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Agreement';

    agreementForm: FormGroup;
    agreement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AgreementService,
        private fb: FormBuilder
) {
        super(http);
        this.agreementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  agreementNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      Customer: ['', ],
      ProductOffering: ['', ],
      AgreementType: ['', ],
      Status: ['', ]
        });
    }

    
    updateAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAgreement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAgreement(params['id']).subscribe(res => {
                this.agreement = res;
            });
        });
    }
}