
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexHealthSystemComponent } from './index.component';
import { HealthSystemService } from '../../../services/HealthSystem.service';

describe('IndexHealthSystemComponent', () => {
  let component: IndexHealthSystemComponent;
  let fixture: ComponentFixture<IndexHealthSystemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexHealthSystemComponent
      ],
      providers: [
        HealthSystemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexHealthSystemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});