
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBIQueryComponent } from './index.component';
import { BIQueryService } from '../../../services/BIQuery.service';

describe('IndexBIQueryComponent', () => {
  let component: IndexBIQueryComponent;
  let fixture: ComponentFixture<IndexBIQueryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBIQueryComponent
      ],
      providers: [
        BIQueryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBIQueryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});