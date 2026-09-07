import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TerminalService } from '../../../services/Terminal.service';
import { Terminal } from '../../../models/Terminal';
import { SubBaseComponent } from '../../Terminal/sub.base.component';

@Component({
    selector: 'app-create-terminal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTerminalComponent extends SubBaseComponent implements OnInit {

    title = 'Add Terminal';

    terminalForm: FormGroup;
    terminal: Terminal;

    constructor( http: HttpClient,
        private terminalService: TerminalService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTerminal(location, Merchant, Type, Status): void {
        this.terminalService
        .addTerminal(location, Merchant, Type, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTerminal']);
            });
    }

    ngOnInit(): void {
    }
}