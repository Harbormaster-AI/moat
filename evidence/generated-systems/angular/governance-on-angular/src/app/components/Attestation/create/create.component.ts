import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AttestationService } from '../../../services/Attestation.service';
import { Attestation } from '../../../models/Attestation';
import { SubBaseComponent } from '../../Attestation/sub.base.component';

@Component({
    selector: 'app-create-attestation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAttestationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Attestation';

    attestationForm: FormGroup;
    attestation: Attestation;

    constructor( http: HttpClient,
        private attestationService: AttestationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.attestationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  statement: ['', Validators.required],
      attestor: ['', Validators.required],
      dateSigned: ['', Validators.required],
      Control: ['', ],
      Policy: ['', ],
      ComplianceProgram: ['', ],
      Result: ['', ]
        });
    }

    
    addAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result): void {
        this.attestationService
        .addAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result)
            .subscribe(() => {
                this.router.navigate(['/indexAttestation']);
            });
    }

    ngOnInit(): void {
    }
}