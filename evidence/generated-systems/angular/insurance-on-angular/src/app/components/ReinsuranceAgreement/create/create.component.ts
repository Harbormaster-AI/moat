import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';
import { ReinsuranceAgreement } from '../../../models/ReinsuranceAgreement';
import { SubBaseComponent } from '../../ReinsuranceAgreement/sub.base.component';

@Component({
    selector: 'app-create-reinsuranceAgreement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateReinsuranceAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Add ReinsuranceAgreement';

    reinsuranceAgreementForm: FormGroup;
    reinsuranceAgreement: ReinsuranceAgreement;

    constructor( http: HttpClient,
        private reinsuranceAgreementService: ReinsuranceAgreementService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.reinsuranceAgreementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  agreementNumber: ['', Validators.required],
      effectivePeriod: ['', Validators.required],
      retention: ['', Validators.required],
      limit: ['', Validators.required],
      cessionPercentage: ['', Validators.required],
      Insurer: ['', ],
      Policies: ['', ],
      ReinsuranceType: ['', ],
      TreatyType: ['', ]
        });
    }

    
    addReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType): void {
        this.reinsuranceAgreementService
        .addReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType)
            .subscribe(() => {
                this.router.navigate(['/indexReinsuranceAgreement']);
            });
    }

    ngOnInit(): void {
    }
}