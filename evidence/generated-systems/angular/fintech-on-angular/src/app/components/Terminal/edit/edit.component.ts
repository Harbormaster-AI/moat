import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TerminalService } from '../../../services/Terminal.service';
import { SubBaseComponent } from '../../Terminal/sub.base.component';


@Component({
    selector: 'app-edit-terminal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTerminalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Terminal';

    terminalForm: FormGroup;
    terminal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TerminalService,
        private fb: FormBuilder
) {
        super(http);
        this.terminalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  location: ['', Validators.required],
      Merchant: ['', ],
      Type: ['', ],
      Status: ['', ]
        });
    }

    
    updateTerminal(location, Merchant, Type, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTerminal(location, Merchant, Type, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTerminal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTerminal(params['id']).subscribe(res => {
                this.terminal = res;
            });
        });
    }
}