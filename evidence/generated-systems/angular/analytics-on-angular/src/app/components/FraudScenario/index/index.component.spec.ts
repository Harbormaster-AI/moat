
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFraudScenarioComponent } from './index.component';
import { FraudScenarioService } from '../../../services/FraudScenario.service';

describe('IndexFraudScenarioComponent', () => {
  let component: IndexFraudScenarioComponent;
  let fixture: ComponentFixture<IndexFraudScenarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFraudScenarioComponent
      ],
      providers: [
        FraudScenarioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFraudScenarioComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});