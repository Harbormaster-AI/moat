
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEmploymentContractComponent } from './create.component';
import { EmploymentContractService } from '../../../services/EmploymentContract.service';
import { Router } from '@angular/router';

describe('CreateEmploymentContractComponent', () => {
  let component: CreateEmploymentContractComponent;
  let fixture: ComponentFixture<CreateEmploymentContractComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEmploymentContractComponent
      ],
      providers: [
        EmploymentContractService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEmploymentContractComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});