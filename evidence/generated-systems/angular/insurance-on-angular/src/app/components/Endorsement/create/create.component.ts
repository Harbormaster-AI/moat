import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EndorsementService } from '../../../services/Endorsement.service';
import { Endorsement } from '../../../models/Endorsement';
import { SubBaseComponent } from '../../Endorsement/sub.base.component';

@Component({
    selector: 'app-create-endorsement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEndorsementComponent extends SubBaseComponent implements OnInit {

    title = 'Add Endorsement';

    endorsementForm: FormGroup;
    endorsement: Endorsement;

    constructor( http: HttpClient,
        private endorsementService: EndorsementService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.endorsementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  endorsementNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      description: ['', Validators.required],
      Policy: ['', ]
        });
    }

    
    addEndorsement(endorsementNumber, effectiveDate, description, Policy): void {
        this.endorsementService
        .addEndorsement(endorsementNumber, effectiveDate, description, Policy)
            .subscribe(() => {
                this.router.navigate(['/indexEndorsement']);
            });
    }

    ngOnInit(): void {
    }
}