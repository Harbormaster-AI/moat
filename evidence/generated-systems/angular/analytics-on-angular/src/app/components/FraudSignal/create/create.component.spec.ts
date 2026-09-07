
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFraudSignalComponent } from './create.component';
import { FraudSignalService } from '../../../services/FraudSignal.service';
import { Router } from '@angular/router';

describe('CreateFraudSignalComponent', () => {
  let component: CreateFraudSignalComponent;
  let fixture: ComponentFixture<CreateFraudSignalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFraudSignalComponent
      ],
      providers: [
        FraudSignalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFraudSignalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});