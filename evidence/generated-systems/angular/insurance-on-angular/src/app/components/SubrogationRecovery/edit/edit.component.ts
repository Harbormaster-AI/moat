import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';
import { SubBaseComponent } from '../../SubrogationRecovery/sub.base.component';


@Component({
    selector: 'app-edit-subrogationRecovery',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSubrogationRecoveryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SubrogationRecovery';

    subrogationRecoveryForm: FormGroup;
    subrogationRecovery: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SubrogationRecoveryService,
        private fb: FormBuilder
) {
        super(http);
        this.subrogationRecoveryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  recoveryReference: ['', Validators.required],
      amount: ['', Validators.required],
      recoveryDate: ['', Validators.required],
      Claim: ['', ],
      Exposure: ['', ],
      Counterparty: ['', ],
      Status: ['', ]
        });
    }

    
    updateSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSubrogationRecovery']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSubrogationRecovery(params['id']).subscribe(res => {
                this.subrogationRecovery = res;
            });
        });
    }
}