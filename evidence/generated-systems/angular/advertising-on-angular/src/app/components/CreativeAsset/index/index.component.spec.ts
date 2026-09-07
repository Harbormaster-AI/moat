
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCreativeAssetComponent } from './index.component';
import { CreativeAssetService } from '../../../services/CreativeAsset.service';

describe('IndexCreativeAssetComponent', () => {
  let component: IndexCreativeAssetComponent;
  let fixture: ComponentFixture<IndexCreativeAssetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCreativeAssetComponent
      ],
      providers: [
        CreativeAssetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCreativeAssetComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});