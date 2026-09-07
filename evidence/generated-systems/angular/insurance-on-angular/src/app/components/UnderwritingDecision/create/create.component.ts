import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';
import { UnderwritingDecision } from '../../../models/UnderwritingDecision';
import { SubBaseComponent } from '../../UnderwritingDecision/sub.base.component';

@Component({
    selector: 'app-create-underwritingDecision',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateUnderwritingDecisionComponent extends SubBaseComponent implements OnInit {

    title = 'Add UnderwritingDecision';

    underwritingDecisionForm: FormGroup;
    underwritingDecision: UnderwritingDecision;

    constructor( http: HttpClient,
        private underwritingDecisionService: UnderwritingDecisionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision): void {
        this.underwritingDecisionService
        .addUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision)
            .subscribe(() => {
                this.router.navigate(['/indexUnderwritingDecision']);
            });
    }

    ngOnInit(): void {
    }
}