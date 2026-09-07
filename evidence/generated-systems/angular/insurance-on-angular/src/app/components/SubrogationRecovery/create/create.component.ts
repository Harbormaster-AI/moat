import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';
import { SubrogationRecovery } from '../../../models/SubrogationRecovery';
import { SubBaseComponent } from '../../SubrogationRecovery/sub.base.component';

@Component({
    selector: 'app-create-subrogationRecovery',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSubrogationRecoveryComponent extends SubBaseComponent implements OnInit {

    title = 'Add SubrogationRecovery';

    subrogationRecoveryForm: FormGroup;
    subrogationRecovery: SubrogationRecovery;

    constructor( http: HttpClient,
        private subrogationRecoveryService: SubrogationRecoveryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status): void {
        this.subrogationRecoveryService
        .addSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status)
            .subscribe(() => {
                this.router.navigate(['/indexSubrogationRecovery']);
            });
    }

    ngOnInit(): void {
    }
}