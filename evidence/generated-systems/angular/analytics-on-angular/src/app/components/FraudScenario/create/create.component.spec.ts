
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFraudScenarioComponent } from './create.component';
import { FraudScenarioService } from '../../../services/FraudScenario.service';
import { Router } from '@angular/router';

describe('CreateFraudScenarioComponent', () => {
  let component: CreateFraudScenarioComponent;
  let fixture: ComponentFixture<CreateFraudScenarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFraudScenarioComponent
      ],
      providers: [
        FraudScenarioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFraudScenarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});