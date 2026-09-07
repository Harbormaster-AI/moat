import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AgreementService } from '../../../services/Agreement.service';
import { Agreement } from '../../../models/Agreement';
import { SubBaseComponent } from '../../Agreement/sub.base.component';

@Component({
    selector: 'app-create-agreement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Add Agreement';

    agreementForm: FormGroup;
    agreement: Agreement;

    constructor( http: HttpClient,
        private agreementService: AgreementService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status): void {
        this.agreementService
        .addAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAgreement']);
            });
    }

    ngOnInit(): void {
    }
}