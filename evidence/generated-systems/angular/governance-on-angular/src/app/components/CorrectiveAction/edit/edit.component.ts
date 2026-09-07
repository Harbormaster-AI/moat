import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CorrectiveActionService } from '../../../services/CorrectiveAction.service';
import { SubBaseComponent } from '../../CorrectiveAction/sub.base.component';


@Component({
    selector: 'app-edit-correctiveAction',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCorrectiveActionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CorrectiveAction';

    correctiveActionForm: FormGroup;
    correctiveAction: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CorrectiveActionService,
        private fb: FormBuilder
) {
        super(http);
        this.correctiveActionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  actionTitle: ['', Validators.required],
      owner: ['', Validators.required],
      targetDate: ['', Validators.required],
      Finding: ['', ],
      Issue: ['', ],
      Status: ['', ]
        });
    }

    
    updateCorrectiveAction(actionTitle, owner, targetDate, Finding, Issue, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCorrectiveAction(actionTitle, owner, targetDate, Finding, Issue, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCorrectiveAction']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCorrectiveAction(params['id']).subscribe(res => {
                this.correctiveAction = res;
            });
        });
    }
}