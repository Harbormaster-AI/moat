import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LegalHoldService } from '../../../services/LegalHold.service';
import { LegalHold } from '../../../models/LegalHold';
import { SubBaseComponent } from '../../LegalHold/sub.base.component';

@Component({
    selector: 'app-create-legalHold',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLegalHoldComponent extends SubBaseComponent implements OnInit {

    title = 'Add LegalHold';

    legalHoldForm: FormGroup;
    legalHold: LegalHold;

    constructor( http: HttpClient,
        private legalHoldService: LegalHoldService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.legalHoldForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      reason: ['', Validators.required],
      issuedDate: ['', Validators.required],
      releaseDate: ['', Validators.required],
      Repositories: ['', ],
      Records: ['', ],
      Matter: ['', ],
      HoldStatus: ['', ]
        });
    }

    
    addLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus): void {
        this.legalHoldService
        .addLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus)
            .subscribe(() => {
                this.router.navigate(['/indexLegalHold']);
            });
    }

    ngOnInit(): void {
    }
}