
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateNonconformanceComponent } from './create.component';
import { NonconformanceService } from '../../../services/Nonconformance.service';
import { Router } from '@angular/router';

describe('CreateNonconformanceComponent', () => {
  let component: CreateNonconformanceComponent;
  let fixture: ComponentFixture<CreateNonconformanceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateNonconformanceComponent
      ],
      providers: [
        NonconformanceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateNonconformanceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});