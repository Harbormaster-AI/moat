import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TerminationService } from '../../../services/Termination.service';
import { SubBaseComponent } from '../../Termination/sub.base.component';


@Component({
    selector: 'app-edit-termination',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTerminationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Termination';

    terminationForm: FormGroup;
    termination: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TerminationService,
        private fb: FormBuilder
) {
        super(http);
        this.terminationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  terminationNumber: ['', Validators.required],
      terminationDate: ['', Validators.required],
      notes: ['', Validators.required],
      eligibleForRehire: ['', Validators.required],
      Employee: ['', ],
      Assignment: ['', ],
      Reason: ['', ],
      Type: ['', ]
        });
    }

    
    updateTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTermination']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTermination(params['id']).subscribe(res => {
                this.termination = res;
            });
        });
    }
}