
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePayrollRunComponent } from './create.component';
import { PayrollRunService } from '../../../services/PayrollRun.service';
import { Router } from '@angular/router';

describe('CreatePayrollRunComponent', () => {
  let component: CreatePayrollRunComponent;
  let fixture: ComponentFixture<CreatePayrollRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePayrollRunComponent
      ],
      providers: [
        PayrollRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePayrollRunComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});