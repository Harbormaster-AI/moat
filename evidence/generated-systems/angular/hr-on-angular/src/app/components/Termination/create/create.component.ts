import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TerminationService } from '../../../services/Termination.service';
import { Termination } from '../../../models/Termination';
import { SubBaseComponent } from '../../Termination/sub.base.component';

@Component({
    selector: 'app-create-termination',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTerminationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Termination';

    terminationForm: FormGroup;
    termination: Termination;

    constructor( http: HttpClient,
        private terminationService: TerminationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type): void {
        this.terminationService
        .addTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type)
            .subscribe(() => {
                this.router.navigate(['/indexTermination']);
            });
    }

    ngOnInit(): void {
    }
}