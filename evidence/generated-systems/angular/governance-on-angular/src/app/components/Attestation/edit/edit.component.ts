import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AttestationService } from '../../../services/Attestation.service';
import { SubBaseComponent } from '../../Attestation/sub.base.component';


@Component({
    selector: 'app-edit-attestation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAttestationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Attestation';

    attestationForm: FormGroup;
    attestation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AttestationService,
        private fb: FormBuilder
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

    
    updateAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAttestation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAttestation(params['id']).subscribe(res => {
                this.attestation = res;
            });
        });
    }
}