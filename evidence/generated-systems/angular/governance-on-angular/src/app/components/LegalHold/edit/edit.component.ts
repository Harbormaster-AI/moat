import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LegalHoldService } from '../../../services/LegalHold.service';
import { SubBaseComponent } from '../../LegalHold/sub.base.component';


@Component({
    selector: 'app-edit-legalHold',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLegalHoldComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LegalHold';

    legalHoldForm: FormGroup;
    legalHold: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LegalHoldService,
        private fb: FormBuilder
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

    
    updateLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLegalHold']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLegalHold(params['id']).subscribe(res => {
                this.legalHold = res;
            });
        });
    }
}