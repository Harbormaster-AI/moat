
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCorrectiveActionComponent } from './index.component';
import { CorrectiveActionService } from '../../../services/CorrectiveAction.service';

describe('IndexCorrectiveActionComponent', () => {
  let component: IndexCorrectiveActionComponent;
  let fixture: ComponentFixture<IndexCorrectiveActionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCorrectiveActionComponent
      ],
      providers: [
        CorrectiveActionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCorrectiveActionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});