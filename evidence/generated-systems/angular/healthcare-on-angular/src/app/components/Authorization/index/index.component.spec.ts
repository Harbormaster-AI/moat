
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAuthorizationComponent } from './index.component';
import { AuthorizationService } from '../../../services/Authorization.service';

describe('IndexAuthorizationComponent', () => {
  let component: IndexAuthorizationComponent;
  let fixture: ComponentFixture<IndexAuthorizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAuthorizationComponent
      ],
      providers: [
        AuthorizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAuthorizationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});