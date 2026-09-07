import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CorrectiveActionService } from '../../../services/CorrectiveAction.service';
import { CorrectiveAction } from '../../../models/CorrectiveAction';
import { SubBaseComponent } from '../../CorrectiveAction/sub.base.component';

@Component({
    selector: 'app-create-correctiveAction',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCorrectiveActionComponent extends SubBaseComponent implements OnInit {

    title = 'Add CorrectiveAction';

    correctiveActionForm: FormGroup;
    correctiveAction: CorrectiveAction;

    constructor( http: HttpClient,
        private correctiveActionService: CorrectiveActionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.correctiveActionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  capaNumber: ['', Validators.required],
      rootCause: ['', Validators.required],
      correctiveAction: ['', Validators.required],
      verificationDate: ['', Validators.required],
      Nonconformance: ['', ],
      Owner: ['', ],
      Status: ['', ]
        });
    }

    
    addCorrectiveAction(capaNumber, rootCause, correctiveAction, verificationDate, Nonconformance, Owner, Status): void {
        this.correctiveActionService
        .addCorrectiveAction(capaNumber, rootCause, correctiveAction, verificationDate, Nonconformance, Owner, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCorrectiveAction']);
            });
    }

    ngOnInit(): void {
    }
}