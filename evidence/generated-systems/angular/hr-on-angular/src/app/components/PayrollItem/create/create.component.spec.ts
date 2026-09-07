
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePayrollItemComponent } from './create.component';
import { PayrollItemService } from '../../../services/PayrollItem.service';
import { Router } from '@angular/router';

describe('CreatePayrollItemComponent', () => {
  let component: CreatePayrollItemComponent;
  let fixture: ComponentFixture<CreatePayrollItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePayrollItemComponent
      ],
      providers: [
        PayrollItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePayrollItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});