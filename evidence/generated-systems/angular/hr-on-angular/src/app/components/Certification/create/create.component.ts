import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CertificationService } from '../../../services/Certification.service';
import { Certification } from '../../../models/Certification';
import { SubBaseComponent } from '../../Certification/sub.base.component';

@Component({
    selector: 'app-create-certification',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCertificationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Certification';

    certificationForm: FormGroup;
    certification: Certification;

    constructor( http: HttpClient,
        private certificationService: CertificationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.certificationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      issuer: ['', Validators.required],
      validFrom: ['', Validators.required],
      validTo: ['', Validators.required],
      credentialId: ['', Validators.required],
      Employee: ['', ],
      Course: ['', ]
        });
    }

    
    addCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course): void {
        this.certificationService
        .addCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course)
            .subscribe(() => {
                this.router.navigate(['/indexCertification']);
            });
    }

    ngOnInit(): void {
    }
}