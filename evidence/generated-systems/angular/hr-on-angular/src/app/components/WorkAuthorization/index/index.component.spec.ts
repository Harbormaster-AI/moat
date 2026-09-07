
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWorkAuthorizationComponent } from './index.component';
import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';

describe('IndexWorkAuthorizationComponent', () => {
  let component: IndexWorkAuthorizationComponent;
  let fixture: ComponentFixture<IndexWorkAuthorizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWorkAuthorizationComponent
      ],
      providers: [
        WorkAuthorizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWorkAuthorizationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});