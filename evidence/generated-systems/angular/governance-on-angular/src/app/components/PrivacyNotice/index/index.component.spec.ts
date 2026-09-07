
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPrivacyNoticeComponent } from './index.component';
import { PrivacyNoticeService } from '../../../services/PrivacyNotice.service';

describe('IndexPrivacyNoticeComponent', () => {
  let component: IndexPrivacyNoticeComponent;
  let fixture: ComponentFixture<IndexPrivacyNoticeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPrivacyNoticeComponent
      ],
      providers: [
        PrivacyNoticeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPrivacyNoticeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});