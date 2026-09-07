
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSubrogationRecoveryComponent } from './index.component';
import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';

describe('IndexSubrogationRecoveryComponent', () => {
  let component: IndexSubrogationRecoveryComponent;
  let fixture: ComponentFixture<IndexSubrogationRecoveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSubrogationRecoveryComponent
      ],
      providers: [
        SubrogationRecoveryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSubrogationRecoveryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});