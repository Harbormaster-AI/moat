
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFraudSignalComponent } from './index.component';
import { FraudSignalService } from '../../../services/FraudSignal.service';

describe('IndexFraudSignalComponent', () => {
  let component: IndexFraudSignalComponent;
  let fixture: ComponentFixture<IndexFraudSignalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFraudSignalComponent
      ],
      providers: [
        FraudSignalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFraudSignalComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});