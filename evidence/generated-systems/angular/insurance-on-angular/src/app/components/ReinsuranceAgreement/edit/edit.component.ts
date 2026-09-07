import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';
import { SubBaseComponent } from '../../ReinsuranceAgreement/sub.base.component';


@Component({
    selector: 'app-edit-reinsuranceAgreement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditReinsuranceAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ReinsuranceAgreement';

    reinsuranceAgreementForm: FormGroup;
    reinsuranceAgreement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ReinsuranceAgreementService,
        private fb: FormBuilder
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

    
    updateReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexReinsuranceAgreement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getReinsuranceAgreement(params['id']).subscribe(res => {
                this.reinsuranceAgreement = res;
            });
        });
    }
}