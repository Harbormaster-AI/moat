import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';
import { SubBaseComponent } from '../../UnderwritingDecision/sub.base.component';


@Component({
    selector: 'app-edit-underwritingDecision',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditUnderwritingDecisionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit UnderwritingDecision';

    underwritingDecisionForm: FormGroup;
    underwritingDecision: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: UnderwritingDecisionService,
        private fb: FormBuilder
) {
        super(http);
        this.underwritingDecisionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  notes: ['', Validators.required],
      decisionDate: ['', Validators.required],
      Quote: ['', ],
      Underwriter: ['', ],
      Decision: ['', ]
        });
    }

    
    updateUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision): void {
        this.route.params.subscribe((params) => {

                        this.service.updateUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexUnderwritingDecision']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getUnderwritingDecision(params['id']).subscribe(res => {
                this.underwritingDecision = res;
            });
        });
    }
}