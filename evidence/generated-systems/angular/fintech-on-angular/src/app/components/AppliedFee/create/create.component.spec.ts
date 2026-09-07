
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAppliedFeeComponent } from './create.component';
import { AppliedFeeService } from '../../../services/AppliedFee.service';
import { Router } from '@angular/router';

describe('CreateAppliedFeeComponent', () => {
  let component: CreateAppliedFeeComponent;
  let fixture: ComponentFixture<CreateAppliedFeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAppliedFeeComponent
      ],
      providers: [
        AppliedFeeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAppliedFeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});